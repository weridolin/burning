interface UserProfile {
    name: string,
    avatar: string,
    email: string,
    phone: string,
    introduction: string,
    roles: string[],
    permissions: string[],
    id: string
    age: number
}
interface AuthToken {
    access_token: string,
    refresh_token: string
}

export function setUserProfile(profile: UserProfile) {
    uni.setStorageSync('profile', JSON.stringify(profile))
}

export function getUserProfile(): UserProfile | null {
    const profile = uni.getStorageSync('profile')
    if (profile) {
        return JSON.parse(profile)
    }
    return null
}

export function setToken(token: AuthToken) {
    uni.setStorageSync('token', JSON.stringify(token))
}

export function getToken(): AuthToken | null {
    const token = uni.getStorageSync('token')
    if (token) {
        return JSON.parse(token)
    }
    return null
}

export function clearToken(){
    uni.removeStorageSync('token')
}

export function isLogin(): boolean {
    return !!getToken()
}
