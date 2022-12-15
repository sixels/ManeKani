use std::{
    io,
    path::{Path, PathBuf},
};

use bytes::Bytes;
use tokio::{fs::File, io::AsyncReadExt};
use uuid::Uuid;

pub struct RequestQuery {
    pub category: String,
    pub name: String,
}

pub struct RequestCreate {
    pub file: Written,
}

pub struct RequestDelete {
    pub name: String,
}

impl RequestQuery {
    pub(crate) fn as_s3_key(&self) -> String {
        s3_key(&self.category, &self.name)
    }
}

impl RequestCreate {
    #[must_use]
    pub fn new(file: Written) -> Self {
        Self { file }
    }

    pub(crate) fn as_s3_key(&self) -> String {
        s3_key(self.file.category(), self.file.name())
    }
}

// impl DeleteFile {
//     pub(crate) fn as_s3_key(&self) -> String {
//         s3_key(&self.category, &self.name)
//     }
// }

fn s3_key(category: &str, name: &str) -> String {
    format!("{}/{}", category, name)
}

pub struct Wrapper {
    path: PathBuf,
    handle: File,
}

pub struct Written {
    path: PathBuf,
    name: String,
    category: String,
    size: u64,
}

impl Wrapper {
    /// Create a new `File` wrapper, used to create temporary files
    ///
    /// # Errors
    ///
    /// This function will return an error if:
    /// - It could not create a temporary directory under /tmp
    /// - The file could not be created in the temporary directory
    /// - An IO error occurred
    pub async fn create() -> io::Result<Self> {
        let base = PathBuf::from("/tmp/manekani/");

        tokio::fs::create_dir_all(&base).await?;

        let path = base.join(Uuid::new_v4().to_string());
        let file = File::create(&path).await?;

        Ok(Self { handle: file, path })
    }

    /// Returns a reference to the path of this [`FileWrapper`].
    pub fn path(&self) -> &Path {
        self.path.as_path()
    }

    /// Finish writing the file, returning a [`WrittenFile`]
    pub async fn finish<C: Into<String>, N: Into<String>>(self, category: C, name: N) -> Written {
        let path = self.path;
        let size = self
            .handle
            .metadata()
            .await
            .map(|meta| meta.len())
            .unwrap_or(1024);

        drop(self.handle);

        Written {
            size,
            category: category.into(),
            name: name.into(),
            path,
        }
    }
}

impl Written {
    /// Read the file and return its contents.
    ///
    /// # Errors
    ///
    /// This function will return an error if the file was not found (or the
    /// user is not allowed to read the file) or another IO error occured
    #[allow(clippy::cast_possible_truncation)]
    pub async fn read_all(&self) -> io::Result<Bytes> {
        let mut file = tokio::fs::File::open(&self.path).await?;

        let mut bytes = Vec::with_capacity(self.size as usize);
        file.read_to_end(&mut bytes).await?;

        Ok(Bytes::from(bytes))
    }

    /// Returns a reference to the name of this [`WrittenFile`].
    #[must_use]
    pub fn name(&self) -> &str {
        self.name.as_ref()
    }

    /// Returns a reference to the category of this [`WrittenFile`].
    #[must_use]
    pub fn category(&self) -> &str {
        self.category.as_ref()
    }
}

impl From<Written> for RequestCreate {
    fn from(val: Written) -> Self {
        RequestCreate { file: val }
    }
}

impl std::ops::Deref for Wrapper {
    type Target = File;

    fn deref(&self) -> &Self::Target {
        &self.handle
    }
}

impl std::ops::DerefMut for Wrapper {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.handle
    }
}

impl Drop for Written {
    fn drop(&mut self) {
        std::fs::remove_file(&self.path).ok();
    }
}
