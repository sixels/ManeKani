/** @type {import('@remix-run/dev').AppConfig} */
export default {
  ignoredRouteFiles: ['**/.*'],
  browserNodeBuiltinsPolyfill: {
    modules: {
      assert: true,
      crypto: true,
      path: true,
      util: true,
      fs: true,
      constants: true,
      buffer: true,
      punycode: true,
      querystring: true,
      timers: true,
      https: true,
      url: true,
      stream: true,
      events: true,
      http: true,
      child_process: true,
      zlib: true,
      os: true,
    },
  },
};
