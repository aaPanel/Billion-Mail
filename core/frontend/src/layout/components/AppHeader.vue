<template>
	<n-layout-header ref="headerRef" :style="{ top: `${top}px` }">
		<div class="header-left">
			<n-button class="icon-btn" :bordered="false" @click="handleCollapse">
				<i class="icon" :class="isCollapse ? 'i-mdi-menu-close' : 'i-mdi-menu-open'"></i>
			</n-button>
		</div>

		<div class="header-right">
			<n-dropdown v-if="langOptions.length > 0" :options="langOptions" @select="handleLangAction">
				<n-button class="icon-btn" :bordered="false">
					<i class="icon i-mdi-language text-20px"></i>
				</n-button>
			</n-dropdown>
			<n-dropdown :options="userOptions" @select="handleUserAction">
				<n-button class="icon-btn" :bordered="false">
					<i class="icon i-mdi-user-outline"></i>
				</n-button>
			</n-dropdown>
		</div>
	</n-layout-header>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { DropdownOption } from 'naive-ui'
import { useUserStore, useGlobalStore } from '@/store'

defineProps({
	top: {
		type: Number,
		default: 0,
	},
})

const { t } = useI18n()

const userStore = useUserStore()
const globalStore = useGlobalStore()

const { isCollapse, langList } = storeToRefs(globalStore)

const handleCollapse = () => {
	globalStore.setCollapse()
}

const langOptions = ref<DropdownOption[]>([])

const userOptions = ref<DropdownOption[]>([
	{
		label: t('layout.menu.logout'),
		key: 'logout',
	},
])

const handleLangAction = async (key: string) => {
	await globalStore.setLang(key)
	window.location.reload()
}

const handleUserAction = (key: string) => {
	switch (key) {
		case 'logout':
			userStore.logout()
			break
	}
}

const getLangOptions = async () => {
	langOptions.value = langList.value.map(item => {
		return {
			label: item.cn,
			key: item.name,
		}
	})
}

onMounted(() => {
	getLangOptions()
})
</script>

<style lang="scss" scoped>
.n-layout-header {
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	display: flex;
	justify-content: space-between;
	align-items: center;
	height: 48px;
	padding: 0 12px;
	background-color: #fff;
	box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
	z-index: 1000;
}

.header-left,
.header-right {
	display: flex;
	align-items: center;
	gap: 8px;
}

.icon-btn {
	--n-width: 42px;
	--n-height: 48px;
	--n-padding: 0;
	--n-font-size: 22px;
	--n-text-color: #5b6b79;
	--n-ripple-color: none;
}
</style>
