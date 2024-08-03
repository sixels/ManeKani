mod ory;

use ory::OryClient;
use serde::Deserialize;

pub use ory::AuthOptions;
use uuid::Uuid;

#[derive(Debug, Clone)]
pub struct SsoSession {
    pub id: String,
    pub traits: Option<SsoSessionTraits>,
    pub metadata_public: Option<SsoSessionMetadata>,
}

#[derive(Debug, Deserialize, Clone)]
pub struct SsoSessionTraits {
    pub email: String,
}

#[derive(Debug, Deserialize, Clone)]
pub struct SsoSessionMetadata {
    pub id: Uuid,
}

#[derive(Debug)]
pub struct AuthManager {
    pub(crate) ory: OryClient,
}

impl AuthManager {
    pub async fn new(options: AuthOptions) -> Result<Self, reqwest::Error> {
        dbg!(reqwest::get(&options.base_url).await?.status());
        let client = OryClient::new(options)?;

        Ok(Self { ory: client })
    }

    pub async fn get_sso_session(&self, session_token: &str) -> Option<SsoSession> {
        let session = match self.ory.frontend_api().to_session(session_token).await {
            Ok(session) => session,
            Err(e) => {
                tracing::error!("failed to get session from cookies: {:?}", e);
                return None;
            }
        };

        let Some(idt) = session.identity else {
            tracing::error!("session has no identity: {:?}", session);
            return None;
        };

        let traits = idt
            .traits
            .and_then(|traits| serde_json::from_value::<SsoSessionTraits>(traits).ok());

        let metadata_public = idt
            .metadata_public
            .and_then(|metadata| serde_json::from_value::<SsoSessionMetadata>(metadata).ok());

        Some(SsoSession {
            id: idt.id,
            traits,
            metadata_public,
        })
    }

    pub fn browser_login_url(&self, return_url: &str) -> String {
        self.ory.frontend_api().browser_login_url(return_url)
    }

    pub async fn is_sign_up_flow_valid(&self, flow_id: &str, cookies: &str) -> bool {
        match self
            .ory
            .frontend_api()
            .get_registration_flow(flow_id, cookies)
            .await
        {
            Ok(_) => true,
            Err(e) => {
                tracing::error!("failed to get registration flow: {:?}", e);
                false
            }
        }
    }
}
