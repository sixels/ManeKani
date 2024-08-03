use actix_web::web;
use manekani_db::Database;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::adapter::user::{self, CreateUserRequest, CreateUserResponse};

#[derive(Debug, Deserialize, Serialize)]
pub struct OnSignUpPayload {
    pub flow: SignUpFlow,
    pub identity: SignUpIdentity,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct SignUpFlow {
    pub id: String,
}
#[derive(Debug, Deserialize, Serialize)]
pub struct SignUpIdentity {
    pub traits: SignUpTraits,
    pub metadata_public: Option<SignUpMetadataPublic>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct SignUpTraits {
    pub email: String,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct SignUpMetadataPublic {
    pub id: Uuid,
}

pub async fn on_sign_up(
    db: web::Data<Database>,
    web::Json(ctx): web::Json<OnSignUpPayload>,
) -> Result<web::Json<OnSignUpPayload>, actix_web::Error> {
    // TODO: check if request is authorized through api key

    let email = &ctx.identity.traits.email;

    dbg!(&ctx);

    match user::create_user(
        &db,
        CreateUserRequest {
            email: email.clone(),
            ..Default::default()
        },
    )
    .await
    {
        Ok(CreateUserResponse(user)) => Ok(web::Json(OnSignUpPayload {
            flow: ctx.flow,
            identity: SignUpIdentity {
                metadata_public: Some(SignUpMetadataPublic { id: user.id }),
                ..ctx.identity
            },
        })),
        Err(e) => match e {
            user::error::CreateUserError::EmailTaken => {
                tracing::info!("user already exists: {}", email);
                return Err(actix_web::error::ErrorConflict(serde_json::json!({
                    "messages": [
                        {
                          "instance_ptr": "#/traits/email",
                          "messages": [
                            {
                              "id": 123,
                              "text": "email already used",
                              "type": "validation",
                              "context": {
                                "value": email,
                              },
                            },
                          ],
                        },
                      ],
                })));
            }
            _ => todo!(),
        },
    }
}
