import EmailPasswordReact from "supertokens-auth-react/recipe/emailpassword";
import SessionReact, {
  BooleanClaim,
} from "supertokens-auth-react/recipe/session";
import { appInfo } from "./appInfo";
import Router from "next/router";

export const frontendConfig = () => {
  return {
    appInfo,
    recipeList: [
      EmailPasswordReact.init({
        useShadowDom: false,
        signInAndUpFeature: {
          signUpForm: {
            formFields: [
              {
                id: "username",
                label: "Username",
                placeholder: "jonh.doe",
                validate: async (username: string) => {
                  const rules = [
                    {
                      match: /^[a-zA-Z]/,
                      message: "Username must start with a letter",
                    },
                    {
                      match: /^[a-zA-Z0-9\.\-_]+$/,
                      message:
                        "Username must contain only letters, numbers, period, '-' and '_'",
                    },
                    {
                      match: (s: string) => s.length > 3 && s.length <= 20,
                      message: "Username length should have 4 to 20 characters",
                    },
                    {
                      match: /^(?!(.*[\.\-_]{2})).*/,
                      message: "Username must have no consecutive punctuation",
                    },
                    {
                      match: /^(?!(.*[\.\-_]$)).+/,
                      message: "Username must end with a letter",
                    },
                  ];

                  for (const rule of rules) {
                    let valid = false;

                    if (rule.match instanceof RegExp) {
                      valid = rule.match.test(username);
                    } else {
                      valid = rule.match(username);
                    }

                    if (!valid) {
                      return rule.message;
                    }
                  }

                  return undefined;
                },
              },
            ],
          },
        },
      }),
      SessionReact.init(),
    ],
    windowHandler: (oI: any) => {
      return {
        ...oI,
        location: {
          ...oI.location,
          setHref: (href: string) => {
            Router.push(href);
          },
        },
      };
    },
  };
};
