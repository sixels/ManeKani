use ory_client::{
    apis::{
        configuration::Configuration,
        frontend_api::{self, GetRegistrationFlowError, ToSessionError},
        Error as OryError,
    },
    models::{RegistrationFlow, Session},
};

#[derive(Debug)]
pub struct OryClient {
    config: Configuration,
    proxy_url: Option<String>,
}

pub struct AuthOptions {
    pub base_url: String,
    pub proxy_url: Option<String>,
}

impl OryClient {
    pub fn new(options: AuthOptions) -> Result<Self, reqwest::Error> {
        let client = {
            let builder = reqwest::Client::builder();

            // let builder = if let Some(proxy_url) = &options.proxy_url {
            //     reqwest::Proxy::all(proxy_url).map(|proxy| builder.proxy(proxy))?
            // } else {
            //     builder
            // };

            builder.build()?
        };

        Ok(Self {
            config: Configuration {
                base_path: options.base_url,
                client,
                ..Default::default()
            },
            proxy_url: options.proxy_url,
        })
    }

    pub fn frontend_api(&self) -> FrontendApi {
        FrontendApi(&self.config, self.proxy_url.as_deref())
    }
}

pub struct FrontendApi<'cfg>(&'cfg Configuration, Option<&'cfg str>);

impl<'cfg> FrontendApi<'cfg> {
    pub async fn to_session(
        &self,
        session_token: &str,
    ) -> Result<Session, OryError<ToSessionError>> {
        frontend_api::to_session(&self.0, None, Some(session_token), None).await
    }

    pub fn browser_login_url(&self, return_url: &str) -> String {
        format!(
            "{}/self-service/login/browser?return_to={}",
            self.1.unwrap_or(&self.0.base_path),
            return_url
        )
    }

    pub async fn get_registration_flow(
        &self,
        flow_id: &str,
        cookies: &str,
    ) -> Result<RegistrationFlow, OryError<GetRegistrationFlowError>> {
        frontend_api::get_registration_flow(&self.0, flow_id, Some(cookies)).await
    }
}
