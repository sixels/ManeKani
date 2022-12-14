use std::{
    io,
    path::{Path, PathBuf},
    str::FromStr,
};

use bytes::Bytes;
use tokio::{fs::File, io::AsyncReadExt};
use uuid::Uuid;

pub struct QueryFile {
    pub category: String,
    pub name: String,
}

pub struct CreateFile {
    pub file: WrittenFile,
}

pub struct DeleteFile {
    pub name: String,
}

impl QueryFile {
    pub(crate) fn as_s3_key(&self) -> String {
        s3_key(&self.category, &self.name)
    }
}

impl CreateFile {
    pub fn new(file: WrittenFile) -> Self {
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

pub struct FileWrapper {
    path: PathBuf,
    handle: File,
}

pub struct WrittenFile {
    path: PathBuf,
    name: String,
    category: String,
    size: u64,
}

impl FileWrapper {
    pub async fn create() -> io::Result<Self> {
        let base = PathBuf::from_str("/tmp/manekani/").unwrap();

        tokio::fs::create_dir_all(&base).await?;

        let path = base.join(Uuid::new_v4().to_string());
        let file = File::create(&path).await?;

        Ok(Self { handle: file, path })
    }

    /// Returns a reference to the path of this [`FileWrapper`].
    pub fn path(&self) -> &Path {
        self.path.as_path()
    }

    pub async fn finish<C: Into<String>, N: Into<String>>(
        self,
        category: C,
        name: N,
    ) -> WrittenFile {
        let path = self.path;
        let size = self
            .handle
            .metadata()
            .await
            .map(|meta| meta.len())
            .unwrap_or(1024);

        drop(self.handle);

        WrittenFile {
            size,
            category: category.into(),
            name: name.into(),
            path,
        }
    }
}

impl WrittenFile {
    pub async fn read_all(&self) -> io::Result<Bytes> {
        let mut file = tokio::fs::File::open(&self.path).await?;

        let mut bytes = Vec::with_capacity(self.size as usize);
        file.read_to_end(&mut bytes).await?;

        Ok(Bytes::from(bytes))
    }

    /// Returns a reference to the name of this [`WrittenFile`].
    pub fn name(&self) -> &str {
        self.name.as_ref()
    }

    /// Returns a reference to the category of this [`WrittenFile`].
    pub fn category(&self) -> &str {
        self.category.as_ref()
    }
}

impl From<WrittenFile> for CreateFile {
    fn from(val: WrittenFile) -> Self {
        CreateFile { file: val }
    }
}

impl std::ops::Deref for FileWrapper {
    type Target = File;

    fn deref(&self) -> &Self::Target {
        &self.handle
    }
}

impl std::ops::DerefMut for FileWrapper {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.handle
    }
}

impl Drop for WrittenFile {
    fn drop(&mut self) {
        std::fs::remove_file(&self.path).ok();
    }
}
