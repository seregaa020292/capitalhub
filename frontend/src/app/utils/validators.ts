import { RuleItem, Rules } from 'async-validator'

interface IElementRule extends RuleItem {
  trigger?: string | string[]
}

interface IValidatorList {
  [index: string]: Array<IElementRule>
}

export const authValidator: IValidatorList = {
  email: [
    { required: true, message: 'Введите Email', trigger: 'blur' },
    {
      type: 'email',
      message: 'Введите правильный Email адрес',
      trigger: ['blur', 'change'],
    },
  ],
  name: [{ required: true, message: 'Введите имя', trigger: 'blur' }],
  password: [
    { required: true, message: 'Введите пароль', trigger: 'blur' },
    { min: 6, message: 'Длина пароля не менее 6 символов', trigger: 'blur' },
  ],
  agree: [
    {
      required: true,
      validator: (rule: Rules, value: boolean) => value,
      message: 'Пожалуйста, согласитесь с условиями.',
      trigger: 'change',
    },
  ],
}

export const assetValidator: IValidatorList = {
  marketId: [{ required: true, message: 'Введите название', trigger: 'blur' }],
  notationAt: [{ required: true, message: 'Введите дату и время', trigger: 'blur' }],
  amount: [{ required: true, message: 'Введите стоимость', trigger: 'blur' }],
  quantity: [{ required: true, message: 'Введите количество', trigger: 'blur' }],
}
