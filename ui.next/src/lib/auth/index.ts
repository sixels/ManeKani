import { Configuration, FrontendApi, Identity } from "@ory/client";
// import { edgeConfig } from "@ory/integrations/next";

export const ORY_SDK_URL = "http://kratos:4433";
export const ORY_BROWSER_URL = "http://127.0.0.1:4433";

export const ory = new FrontendApi(
	new Configuration({
		basePath: ORY_SDK_URL,
		baseOptions: {
			withCredentials: true,
		},
	}),
);

export const getLoginURL = (returnTo: string): string => {
	return `${ORY_BROWSER_URL}/self-service/login/browser?return_to=${returnTo}`;
};
