pub mod file;
pub mod message;

use std::{
    fs::File,
    io::{Error as IoError, ErrorKind as IoErrorKind, Result as IoResult, Write},
    path::{Path, PathBuf},
};

use bytes::Bytes;
use crossbeam_channel::Sender;
use uuid::Uuid;

use self::message::{Message, MessageKind};

pub struct FilesystemOverlay {}

impl FilesystemOverlay {
    pub fn start<P: AsRef<Path>>(base_directory: P) -> Sender<Message> {
        let (sender, receiver) = crossbeam_channel::unbounded::<Message>();

        let base = PathBuf::from(base_directory.as_ref());
        // let this = Self::new(receiver, dir);

        std::thread::spawn(move || {
            while let Ok(message) = receiver.recv() {
                let res = match message.kind() {
                    MessageKind::QueryFile(filename) => {
                        let path = base.join(filename);
                        if query_file(&path) {
                            Ok(path)
                        } else {
                            Err(IoError::from(IoErrorKind::NotFound))
                        }
                    }

                    MessageKind::CreateFile { filetype, contents } => {
                        let filename = Uuid::new_v4().to_string();
                        let filename = format!("{filename}.{filetype}");
                        let path = base.join(filename);
                        create_file(&path, contents).map(|_| path)
                    }

                    MessageKind::DeleteFile(_filename) => {
                        todo!()
                    }

                    MessageKind::Stop => break,
                };

                if message
                    .reply(res.map(|p| p.to_str().unwrap().to_owned()))
                    .is_err()
                {
                    tracing::warn!("Filesystem overlay failed to send a reply to the request.")
                };
            }
        });

        sender
    }
}

fn create_file<P: AsRef<Path>>(path: P, contents: &Bytes) -> IoResult<()> {
    let path = path.as_ref();
    if query_file(path) {
        return Err(IoError::from(IoErrorKind::AlreadyExists));
    }

    let mut file = File::create(path)?;
    file.write_all(contents)?;

    Ok(())
}

fn query_file<P: AsRef<Path>>(path: P) -> bool {
    let path = path.as_ref();
    path.exists()
}
