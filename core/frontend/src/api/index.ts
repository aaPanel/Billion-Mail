import axios from 'axios'
import router from '@/router'
import { useUserStore } from '@/store'
import { apiUrlPrefix, isObject } from '@/utils'
import { Message } from '@/utils/message'

interface FetchOptions {
	prefix: string
}

const instance = axios.create({
	baseURL: apiUrlPrefix,
	timeout: 250000,
	headers: {
		'Content-Type': 'application/json',
	},
	fetchOptions: {
		prefix: '/api',
		successMessage: false,
	},
})

const whitePathList = ['/login']

// 请求拦截器 增加token
instance.interceptors.request.use(config => {
	const { headers, url } = config
	if (!whitePathList.includes(url || '')) {
		const userStore = useUserStore()
		headers.Authorization = `Bearer ${userStore.login.token}`
	}
	return config
})

// 请求拦截器 处理请求前缀
instance.interceptors.request.use(config => {
	const { fetchOptions } = config
	if (isObject<FetchOptions>(fetchOptions)) {
		config.url = `${fetchOptions.prefix}` + config.url
	}
	return config
})

// 成功、失败处理
instance.interceptors.response.use(
	response => {
		const { fetchOptions } = response.config
		const { code, data, msg, success } = response.data || {}
		if (code === 0 && success) {
			if (fetchOptions?.successMessage) {
				Message.success(msg)
			}
			return Promise.resolve(data)
		}
		if (!success && msg) {
			Message.error(msg)
		}
		if (code === 401) {
			router.push('/login')
		}
		return Promise.reject(response.data)
	},
	error => {
		return Promise.reject(error)
	}
)

export { instance }
