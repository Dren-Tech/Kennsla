import { Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";
import { MiddlewareInterface } from "./MiddlewareInterface.ts";

const RequestLoggerMiddleware: MiddlewareInterface = async function(ctx: Context, next: CallableFunction): Promise<void> {
    await next();

    console.log(`${ctx.request.method}: ${ctx.request.url} ${ctx.response.status} (${ctx.response.headers.get("X-Response-Time")})`);
}

export default RequestLoggerMiddleware;