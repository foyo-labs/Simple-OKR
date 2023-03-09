const isDev = process.env.NODE_ENV === 'development'
const apiUrl = isDev ? 'http://localhost:8080/api':'http://www4.ckmro.com/api'

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  publicRuntimeConfig: {
		apiUrl: apiUrl,
	}
}

module.exports = nextConfig
