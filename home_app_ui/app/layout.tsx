import type { Metadata } from "next";

// These styles apply to every route in the application
import "../styles/global.css";
import NavigationBar from "./components/navigation/navigationBar";

export const metadata: Metadata = {
  title: "My Home Apps",
  description: "App for my home personal managment",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className="bg-body">
        <div className="flex min-h-screen">
          <NavigationBar></NavigationBar>
          <main className="flex-1 p-4">{children}</main>
        </div>
      </body>
    </html>
  );
}
