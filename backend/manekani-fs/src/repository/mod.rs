use std::{
    io,
    path::{Path, PathBuf},
};

use crossbeam_channel::Sender;
use manekani_types::repository::{InsertError, QueryError, RepoInsertable, RepoQueryable};

use crate::fs::{file::File, message::Message, FilesystemOverlay};

pub struct FilesystemRepo {
    fs_messages: Sender<Message>,
}

impl FilesystemRepo {
    pub fn new<P: AsRef<Path>>(storage_location: P) -> Self {
        let base_location = PathBuf::from(storage_location.as_ref());
        let sender = FilesystemOverlay::start(base_location);

        Self {
            fs_messages: sender,
        }
    }
}

impl Drop for FilesystemRepo {
    fn drop(&mut self) {
        let (msg, _) = Message::stop();
        // there is nothing more we can do
        self.fs_messages.send(msg).ok();
    }
}

#[async_trait::async_trait]
impl RepoQueryable<String, String> for FilesystemRepo {
    async fn query(&self, filename: String) -> Result<String, QueryError> {
        let (msg, receiver) = Message::query_file(filename);
        if self.fs_messages.send(msg).is_err() {
            return Err(QueryError::Unknown);
        };

        let Ok(result) = receiver.await else {
            return Err(QueryError::Unknown);
        };

        match result {
            Ok(filename) => Ok(filename),
            Err(error) => match error.kind() {
                io::ErrorKind::NotFound => Err(QueryError::NotFound),
                _ => Err(QueryError::Unknown),
            },
        }
    }
}

#[async_trait::async_trait]
impl RepoInsertable<File, String> for FilesystemRepo {
    async fn insert(&self, file: File) -> Result<String, InsertError> {
        let (msg, receiver) = Message::create_file(file.filetype, file.contents);
        if self.fs_messages.send(msg).is_err() {
            return Err(InsertError::Unknown);
        };
        let Ok(result) = receiver.await else {
            return Err(InsertError::Unknown);
        };

        match result {
            Ok(filename) => Ok(filename),
            Err(error) => match error.kind() {
                io::ErrorKind::AlreadyExists => Err(InsertError::Conflict),
                _ => Err(InsertError::Unknown),
            },
        }
    }
}
