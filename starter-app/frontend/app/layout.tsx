// The root layout wraps every page. Put shared things here (nav bar, fonts).
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Todo Starter",
  description: "Next.js + Go learning project",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <main>{children}</main>
      </body>
    </html>
  );
}
