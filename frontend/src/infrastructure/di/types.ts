export default {
  IStoreRoot: Symbol.for('IStoreRoot'),

  IAuthPresenter: Symbol.for('IAuthPresenter'),
  IUserPresenter: Symbol.for('IUserPresenter'),
  IAssetPresenter: Symbol.for('IAssetPresenter'),
  IPortfolioPresenter: Symbol.for('IPortfolioPresenter'),

  IAssetFetchUseCase: Symbol.for('IAssetFetchUseCase'),
  IAssetAddUseCase: Symbol.for('IAssetAddUseCase'),
  IUserFetchUseCase: Symbol.for('IUserFetchUseCase'),
  IMarketSearchUseCase: Symbol.for('IMarketSearchUseCase'),
  IPortfolioFetchUseCase: Symbol.for('IPortfolioFetchUseCase'),
  IPortfoliosFetchUseCase: Symbol.for('IPortfoliosFetchUseCase'),
  IPortfolioAddUseCase: Symbol.for('IPortfolioAddUseCase'),

  IUserRepository: Symbol.for('IUserRepository'),
  IAuthRepository: Symbol.for('IAuthRepository'),
  IAssetRepository: Symbol.for('IAssetRepository'),
  IMarketRepository: Symbol.for('IMarketRepository'),
  IPortfolioRepository: Symbol.for('IPortfolioRepository'),

  IUserClientApi: Symbol.for('IUserClientApi'),
  IAssetClientApi: Symbol.for('IAssetClientApi'),
  IAuthClientApi: Symbol.for('IAuthClientApi'),
  IMarketClientApi: Symbol.for('IMarketClientApi'),
  IPortfolioClientApi: Symbol.for('IPortfolioClientApi'),

  IAuthService: Symbol.for('IAuthService'),
  IStorageService: Symbol.for('IStorageService'),
  IMessageService: Symbol.for('IMessageService'),
  INotifyService: Symbol.for('INotifyService'),
  ILoggerService: Symbol.for('ILoggerService'),
  IRouterService: Symbol.for('IRouterService'),

  IErrorHandler: Symbol.for('IErrorHandler'),
}
