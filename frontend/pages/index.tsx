import React, { useEffect } from "react";
import { useRouter } from "next/router";

const HomePage = () => {
  const router = useRouter();

  useEffect(() => {
    // コンポーネントのマウント時に/loginにリダイレクト
    router.push('/login');
  }, [router]); // 依存配列にrouterを指定

  return <div>Redirecting to login...</div>;
};

export default HomePage;
