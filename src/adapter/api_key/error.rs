use manekani_db::query::api_key;

#[derive(Debug)]
pub enum GetApiKeyError {
    NotFound,
    InternalError(Box<dyn std::error::Error>),
}

#[derive(Debug)]
pub enum GetApiKeysError {
    InternalError,
}

#[derive(Debug)]
pub enum CreateApiKeyError {
    LimitExceeded(usize),
    DuplicateApiKey,
    InternalError,
    // TODO: struct validation
    #[allow(dead_code)]
    ValidationFailed,
}

impl From<api_key::read::GetApiKeyError> for GetApiKeyError {
    fn from(err: api_key::read::GetApiKeyError) -> Self {
        match err {
            api_key::read::GetApiKeyError::NotFound => GetApiKeyError::NotFound,
            api_key::read::GetApiKeyError::DatabaseError(e) => {
                GetApiKeyError::InternalError(e.into())
            }
        }
    }
}

impl From<api_key::read::GetUserApiKeysError> for GetApiKeysError {
    fn from(err: api_key::read::GetUserApiKeysError) -> Self {
        match err {
            api_key::read::GetUserApiKeysError::DatabaseError(_) => GetApiKeysError::InternalError,
        }
    }
}

impl From<api_key::read::CountUserApiKeysError> for CreateApiKeyError {
    fn from(err: api_key::read::CountUserApiKeysError) -> Self {
        match err {
            api_key::read::CountUserApiKeysError::DatabaseError(_) => {
                CreateApiKeyError::InternalError
            }
        }
    }
}

impl From<api_key::write::CreateApiKeyError> for CreateApiKeyError {
    fn from(err: api_key::write::CreateApiKeyError) -> Self {
        match err {
            api_key::write::CreateApiKeyError::DuplicateApiKey => {
                CreateApiKeyError::DuplicateApiKey
            }
            api_key::write::CreateApiKeyError::DatabaseError(_) => CreateApiKeyError::InternalError,
        }
    }
}
