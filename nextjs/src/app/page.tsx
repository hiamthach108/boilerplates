import Link from 'next/link';

import HomeHero from '@/components/features/home/HomeHero';

export default function Home() {
  return (
    <div className="mx-auto flex w-full max-w-5xl flex-col gap-8 px-6 py-12">
      <HomeHero />
      <Link className="text-sm font-medium text-blue-600 hover:text-blue-700" href="/health-check">
        Health check
      </Link>
    </div>
  );
}
