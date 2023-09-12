import defaultConfig, { paths } from './jest.config';

import { pathsToModuleNameMapper } from 'ts-jest';

export default {
  ...defaultConfig,
  moduleNameMapper: pathsToModuleNameMapper(paths, {
    prefix: '<rootDir>/../src',
  }),
  rootDir: './test',
  testRegex: '.e2e-spec\\.ts$',
};
