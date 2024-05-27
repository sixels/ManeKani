use rand::{rngs::OsRng, RngCore};

use crate::util::sha256;

pub struct TokenData {
    pub prefix: String,
    pub hash: String,
    pub key: String,
}

pub fn generate_random_token() -> TokenData {
    let prefix_bytes = {
        let mut random_bytes = [0u8; 4];
        OsRng::default().fill_bytes(&mut random_bytes);
        random_bytes
    };
    let token_uuid = uuid::Uuid::new_v4();
    let token_bytes = token_uuid.as_bytes();

    let hash = sha256::hash(token_bytes.as_ref(), &prefix_bytes);

    let prefix = base16ct::lower::encode_string(&prefix_bytes);
    let token = bs58::encode(token_bytes).into_string();

    let key = format!("{prefix}.{token}");

    TokenData { prefix, hash, key }
}
