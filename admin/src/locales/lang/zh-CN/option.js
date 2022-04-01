import pay from './option/pay'
import oss from './option/oss'
import email from './option/email'
export default {
  ...pay,
  ...oss,
  ...email
}
