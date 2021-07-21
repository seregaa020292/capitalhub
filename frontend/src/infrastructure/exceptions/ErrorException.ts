export abstract class ErrorException extends Error {
  constructor(message?: string) {
    super(message)

    if (Error.captureStackTrace) {
      Error.captureStackTrace(this, this.constructor)
    }
  }

  abstract get result(): unknown
}
