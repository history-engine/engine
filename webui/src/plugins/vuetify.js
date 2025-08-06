/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify } from 'vuetify'

// https://vuetifyjs.com/zh-Hans/components/number-inputs/
import { VNumberInput } from 'vuetify/labs/VNumberInput'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
// https://vuetifyjs.com/en/features/theme/
export default createVuetify({
  theme: {
    defaultTheme: 'light',
  },
  components: {
    VNumberInput,
  },
})
