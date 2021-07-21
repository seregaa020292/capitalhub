export interface BaseHandler<IParam, IResult> {
  handle(params?: IParam): IResult
}
