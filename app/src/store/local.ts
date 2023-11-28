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

