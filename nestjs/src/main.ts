import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import { SwaggerModule, DocumentBuilder } from "@nestjs/swagger";
import { ValidationPipe } from "@nestjs/common";
import * as fs from "fs";
import {
  APP_VERSION,
  HTTP_CORS,
  HTTP_CORS_METHODS,
  HTTP_PORT,
} from "./shared/constants/env.const";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.enableCors({
    origin: HTTP_CORS,
    methods: HTTP_CORS_METHODS,
    preflightContinue: false,
    optionsSuccessStatus: 204,
  });

  // Swagger
  const config = new DocumentBuilder()
    .setTitle("Events Management System")
    .setDescription("Events Management System")
    .setVersion(APP_VERSION)
    .addBearerAuth()
    .build();
  const document = SwaggerModule.createDocument(app, config);

  SwaggerModule.setup("swagger", app, document);

  fs.writeFileSync("./swagger-spec.json", JSON.stringify(document, null, 2));

  SwaggerModule.setup("api/docs", app, document);

  app.useGlobalPipes(new ValidationPipe());

  await app.listen(HTTP_PORT);
}
bootstrap();
