import Link from 'next/link';

export default function Header() {
  return (
    <header className="border-b border-neutral-200">
      <div className="mx-auto flex h-14 w-full max-w-5xl items-center justify-between px-6">
        <Link className="font-semibold text-neutral-950" href="/">
          Dreon Starter
        </Link>
      </div>
    </header>
  );
}
