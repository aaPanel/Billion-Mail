import { createI18n } from 'vue-i18n'
import en from './lang/en.json'
import zh from './lang/zh.json'

const i18n = createI18n({
	legacy: false,
	globalInjection: true,
	locale: 'en',
	fallbackLocale: 'en',
	messages: {
		en,
		zh,
	},
})

// 提供切换语言的方法
export const setLanguage = (locale: string) => {
	i18n.global.locale.value = locale as 'zh' | 'en'
}

export default i18n
