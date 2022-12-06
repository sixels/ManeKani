use std::io::Result as IoResult;

use bytes::Bytes;
use tokio::sync::oneshot::{self, Receiver as OneshotReceiver, Sender as OneshotSender};

pub type MessageReceiver = OneshotReceiver<IoResult<String>>;

pub struct Message {
    sender: OneshotSender<IoResult<String>>,
    kind: MessageKind,
}

impl Message {
    pub fn new(sender: OneshotSender<IoResult<String>>, kind: MessageKind) -> Self {
        Self { sender, kind }
    }

    pub(super) fn reply(self, res: IoResult<String>) -> Result<(), IoResult<String>> {
        self.sender.send(res)
    }

    pub fn kind(&self) -> &MessageKind {
        &self.kind
    }

    pub fn query_file(filename: String) -> (Message, MessageReceiver) {
        let (sender, receiver) = oneshot::channel();

        (
            Message::new(sender, MessageKind::QueryFile(filename)),
            receiver,
        )
    }
    pub fn create_file(filetype: String, contents: Bytes) -> (Message, MessageReceiver) {
        let (sender, receiver) = oneshot::channel();

        (
            Message::new(sender, MessageKind::CreateFile { filetype, contents }),
            receiver,
        )
    }
    pub fn delete_file(filename: String) -> (Message, MessageReceiver) {
        let (sender, receiver) = oneshot::channel();

        (
            Message::new(sender, MessageKind::DeleteFile(filename)),
            receiver,
        )
    }
    pub fn stop() -> (Message, MessageReceiver) {
        let (sender, receiver) = oneshot::channel();
        (Message::new(sender, MessageKind::Stop), receiver)
    }
}

pub enum MessageKind {
    QueryFile(String),
    CreateFile { filetype: String, contents: Bytes },
    DeleteFile(String),
    Stop,
}
