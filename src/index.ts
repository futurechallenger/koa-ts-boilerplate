import Koa from "koa";
import Router from "koa-router";
import BodyParser from "koa-bodyparser";

const app = new Koa();
const router = new Router();
const bodyParser = BodyParser();

app.use(bodyParser);

/**
 * General params for executing a method in a file
 */
router.post("/exec", async ctx => {
  const {request} = ctx
  console.log(request.query);

  const {
    funcName,
    params,
  }=  request.query;
  
  ctx.body = "hello world";
});

app.use(router.routes());

app.listen(55667)