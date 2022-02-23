/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  env: {
    NEXT_PUBLIC_API_URL: process.env.API_URL,
  },
  experimental: {
    outputStandalone: true,
  },
};

module.exports = nextConfig;
