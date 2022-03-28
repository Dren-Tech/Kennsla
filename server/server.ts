import { Application, Router, Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";
import JsonResponseMiddleware from "./src/http/middleware/JsonResponseMiddleware.ts";

const router = new Router();
router.get("/", (ctx: Context) => {
  ctx.response.body = "Hello world!";
});
router.get("/task/{slug}", (ctx: Context) => {
  ctx.response.body = {"hello": "world"};
});

const app = new Application();
app.use(router.routes());
app.use(router.allowedMethods());
app.use(JsonResponseMiddleware);

app.addEventListener(
  "listen",
  (_e) => console.log("Listening on http://localhost:8080"),
);
await app.listen({ port: 8080 });
