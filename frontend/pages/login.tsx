import { useState } from "react";
import axios from "axios";
import { useRouter } from "next/router";
import styles from "./login.module.scss"; // 仮定のスタイルシートのパス

export default function Login() {
  // ログイン用の状態
  const [loginEmail, setLoginEmail] = useState("");
  const [loginPassword, setLoginPassword] = useState("");
  // 新規登録用の状態
  const [name, setName] = useState("");
  const [registerEmail, setRegisterEmail] = useState("");
  const [registerPassword, setRegisterPassword] = useState("");
  const router = useRouter();

  // バックエンドAPIのベースURL
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_BASE_URL;
  
  console.log(process.env)
  // ログイン処理
  const handleSubmitLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      // ログイン処理のAPI呼び出し
      const { data } = await axios.post(`${backendUrl}/users/login`, {
        email: loginEmail,
        password: loginPassword,
      });
      localStorage.setItem("token", data.authtoken); // 認証トークンをローカルストレージに保存
      router.push("/main"); // メインページにリダイレクト
    } catch (error) {
      console.error(error);
      alert("ログインに失敗しました。");
    }
  };

  // 新規登録処理
  const handleSubmitRegistration = async (
    e: React.FormEvent<HTMLFormElement>
  ) => {
    e.preventDefault();
    try {
      // 新規登録処理のAPI呼び出し
      const { data } = await axios.post(`${backendUrl}/users/me`, {
        name,
        email: registerEmail,
        password: registerPassword,
      });
      alert("登録が完了しました。ログインしてください。");
      // 登録後はログインページにリダイレクト等の処理
    } catch (error) {
      console.error(error);
      alert("登録に失敗しました。");
    }
  };

  return (
    <>
      <form onSubmit={handleSubmitLogin} className={styles.loginForm}>
        <h3>ログイン</h3>
        <input
          type="email"
          value={loginEmail}
          onChange={(e) => setLoginEmail(e.target.value)}
          placeholder="Email"
          required
        />
        <input
          type="password"
          value={loginPassword}
          onChange={(e) => setLoginPassword(e.target.value)}
          placeholder="Password"
          required
        />
        <button type="submit">ログイン</button>
      </form>
      <form
        onSubmit={handleSubmitRegistration}
        className={styles.registrationForm}
      >
        <h3>新規登録</h3>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Name"
          required
        />
        <input
          type="email"
          value={registerEmail}
          onChange={(e) => setRegisterEmail(e.target.value)}
          placeholder="Email"
          required
        />
        <input
          type="password"
          value={registerPassword}
          onChange={(e) => setRegisterPassword(e.target.value)}
          placeholder="Password"
          required
        />
        <button type="submit">登録</button>
      </form>
      <img
        src="/cat1.png"
        alt="Description of image"
        className={styles.cardImage}
      />
    </>
  );
}
