import React from "react";
import {
  SessionAuth,
  useSessionContext,
} from "supertokens-auth-react/recipe/session";

export default function AuthRoute(props: React.PropsWithChildren<any>) {
  return (
    <SessionAuth
      overrideGlobalClaimValidators={(globalValidators) => [
        ...globalValidators,
        // TODO: add claim validator for email verification and other stuff
        // [...].validators.[...](),
      ]}
    >
      <InvalidClaimHandler>{props.children}</InvalidClaimHandler>
    </SessionAuth>
  );
}

function InvalidClaimHandler(props: React.PropsWithChildren<any>) {
  let sessionContext = useSessionContext();

  if (sessionContext.loading) {
    return null;
  }

  // check claims
  // if (
  //   sessionContext.invalidClaims.some(
  //     (i) => i.validatorId === TODO
  //   )
  // ) {
  //   // TODO: redirect to username set page
  //   return <div>You cannot access this page because you are not an admin.</div>;
  // }

  // We show the protected route since all claims validators have
  // passed implying that the user is an admin.
  return <div>{props.children}</div>;
}
