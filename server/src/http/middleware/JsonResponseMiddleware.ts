import { Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";
import { MiddlewareInterface } from "./MiddlewareInterface.ts";

const JsonResponseMiddleware: MiddlewareInterface = async function (
  ctx: Context,
  next: CallableFunction,
): Promise<void> {
  if (!ctx.response.headers.has("content-type")) {
    ctx.response.headers.append("content-type", "application/json");
  }

  await next();
};

export default JsonResponseMiddleware;
