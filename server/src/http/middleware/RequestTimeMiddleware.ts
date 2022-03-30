import { Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";
import { MiddlewareInterface } from "./MiddlewareInterface.ts";

const RequestTimeMiddleware: MiddlewareInterface = async function(ctx: Context, next: CallableFunction): Promise<void> {
  const start = Date.now();
  await next();
  const ms = Date.now() - start;
  
  ctx.response.headers.set("X-Response-Time", `${ms}ms`);
}

export default RequestTimeMiddleware;