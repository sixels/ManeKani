local claims = {
  email_verified: true,
} + std.extVar('claims');

{
  identity: {
    traits: {
      [if 'email' in claims && claims.email_verified then 'email' else null]: claims.email,
      display_name: claims.given_name,
      [if 'hd' in claims && claims.email_verified then 'hd' else null]: claims.hd,
    },
  },
}
