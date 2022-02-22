/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: process.env.API_URL + "/:path*",
      },
    ];
  },
  env: {
    NEXT_PUBLIC_API_URL: process.env.API_URL,
  },
};

module.exports = nextConfig;
