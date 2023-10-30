import { AuthModule } from './hooks.auth.module';
import { Module } from '@nestjs/common';

@Module({
  imports: [AuthModule],
  controllers: [],
  providers: [],
  exports: [],
})
export class HooksModule {}
