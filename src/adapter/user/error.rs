use manekani_db::query::user;

#[derive(Debug)]
pub enum CreateUserError {
    EmailTaken,
    InvalidEmail,
    InvalidUsername,
    UsernameTaken,
    InternalError(Box<dyn std::error::Error>),
}

impl From<user::write::CreateUserError> for CreateUserError {
    fn from(err: user::write::CreateUserError) -> Self {
        match err {
            user::write::CreateUserError::DuplicateEmail => CreateUserError::EmailTaken,
            user::write::CreateUserError::DuplicateUsername => CreateUserError::UsernameTaken,
            user::write::CreateUserError::DatabaseError(e) => {
                CreateUserError::InternalError(e.into())
            }
        }
    }
}
