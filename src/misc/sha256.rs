use sha2::{Digest, Sha256};

pub fn hash(data: &[u8], prefix: &[u8]) -> String {
    let mut hasher = Sha256::new();

    hasher.update(prefix);
    hasher.update(data);

    let hash = hasher.finalize();
    let hash_encoded = base16ct::lower::encode_string(&hash);

    hash_encoded
}
