import { Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";
import { MiddlewareInterface } from "./MiddlewareInterface.ts";

const JsonResponseMiddleware: MiddlewareInterface = function(ctx: Context, next: CallableFunction): void {
    if(!ctx.response.headers.has("content-type")){
        ctx.response.headers.append("content-type", "application/json");
    }

    next();
}

export default JsonResponseMiddleware;