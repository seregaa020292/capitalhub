export interface BaseUseCase<IParam, IResult> {
  execute(params?: IParam): IResult
}
