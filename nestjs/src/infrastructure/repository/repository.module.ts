import { Global, Module } from "@nestjs/common";
import { DbService } from "../database/db.service";
import { UserRepository } from "./user.repository";

@Global()
@Module({
  imports: [],
  providers: [DbService, UserRepository],
  exports: [UserRepository],
})
export class RepositoryModule {}
