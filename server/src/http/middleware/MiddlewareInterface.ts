import { Context } from "https://deno.land/x/oak@v10.5.1/mod.ts";

interface MiddlewareInterface {
    (ctx: Context, next: CallableFunction): void;
}

export type {MiddlewareInterface};