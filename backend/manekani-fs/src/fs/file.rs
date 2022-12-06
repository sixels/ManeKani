use bytes::Bytes;

pub struct File {
    pub filetype: String,
    pub contents: Bytes,
}
