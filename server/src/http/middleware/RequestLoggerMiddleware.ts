import { Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";
import { MiddlewareInterface } from "./MiddlewareInterface.ts";

const RequestLoggerMiddleware: MiddlewareInterface = function(ctx: Context, next: CallableFunction): void {
    next();

    console.log(`${ctx.request.method}: ${ctx.request.url} [${ctx.response.status}]`);
}

export default RequestLoggerMiddleware;