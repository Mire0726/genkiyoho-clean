import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import axios, { AxiosRequestConfig } from "axios";
import styles from "./main.module.scss"; // SCSSモジュールのインポート

type Condition = {
  condition_name: string;
  // start_date: string;
  // end_date: string;
  damage_point: number;
  Name : string;
};
export default function Main() {
  const [conditions, setConditions] = useState<Condition[]>([]);
  const [allConditions, setAllConditions] = useState<Condition[]>([]);
  const [errorMessage, setErrorMessage] = useState("");
  const router = useRouter();
  const [genkiHP, setGenkiHP] = useState(null);
  const [todayMassage, setTodayMassage] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login"); // トークンがなければログインページにリダイレクト
    } else {
      fetchConditionsDisplay(token);
      todayPoint(token);
      handlefetchConditions(token);
    }
  }, [router]);
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_BASE_URL;

  const fetchConditionsDisplay = async (token: string) => {
    console.log("Fetching conditions...");
    const options: AxiosRequestConfig = {
      url: `${backendUrl}/users/me/condition/today`,
      method: "GET",
      headers: {
        "x-token": token,
      },
    };
    console.log("Options:", options);

    // エラー処理
    try {
      const response = await axios(options);
      if (Array.isArray(response.data)) {
        const data = response.data;
        setConditions(data);
      } else {
        setConditions([]);
      }
      setErrorMessage("");
    } catch (error) {
      console.error("Error fetching conditions:", error);
      setErrorMessage(
        "情報の取得中にサーバーでエラーが発生しました。しばらくしてから再度試してください。"
      );
    }
    console.log("Conditions after fetching:", conditions); // ログに状態を出力
  };

  const todayPoint = async (token: string) => {
    console.log("Fetching today's point...");
    const options: AxiosRequestConfig = {
      url: `${backendUrl}/users/me/condition/today/point`,
      method: "GET",
      headers: {
        "x-token": token,
      },
    };
    
  try {
    const response = await axios(options);
    if (response.data !== null) {
      const genkiHP = response.data;
      console.log(`Today's Genki HP:`, genkiHP);
      setGenkiHP(genkiHP);
      if (genkiHP <= 50) {
        setTodayMassage("今日はゆっくりめに過ごしましょう");
      }
    } else {
      // 応答がnullの場合は、genkiHPを初期値または空値に設定
      setGenkiHP(null);
    }
    setErrorMessage("");
  } catch (error) {
    console.error("Error fetching today's point:", error);
    setErrorMessage("情報の取得中にサーバーでエラーが発生しました。しばらくしてから再度試してください。");
  }
};

  const handlefetchConditions = async (token: string) => {
    console.log("Fetching conditions...");
    const options: AxiosRequestConfig = {
      url: `${backendUrl}/users/me/condition`,
      method: "GET",
      headers: {
        "x-token": token,
      },
    };

    try {
      const response = await axios(options);
      const data = response.data;

      if (data !==null){
        setAllConditions(data);
        setErrorMessage("");
      } else{
        setAllConditions([]);
      }
    } catch (error) {
      console.error("Error fetching conditions:", error);
      setErrorMessage(
        "情報の取得中にサーバーでエラーが発生しました。しばらくしてから再度試してください。"
      );
    }
  };

  const handleConditionClick = () => {
    router.push("/condition"); // コンディションページへのリダイレクト
  };

  return (
    <div className={styles.mainContainer}>
      <div className={styles.card}>
        <h1>今日の元気予報</h1>
        <h2>
          {genkiHP !== null && <p>{genkiHP}/100</p>} {/* 元気ポイントの表示 */}
        </h2>
        {todayMassage && <p>{todayMassage}</p>} {/* メッセージの表示 */}
        {/* Rest of your code */}
        <img
          src="/girl1.png"
          alt="Description of image"
          className={styles.cardImage}
        />
      </div>
      <div className={styles.cards}>
        <div className={styles.cardmini}>
          <h2>予報詳細：</h2>
          <div className={styles.detail}>
            <ul>
              {conditions.map((condition, index) => (
                <li key={index}>
                  {condition.condition_name}で-{condition.damage_point}pt
                </li>
              ))}
            </ul>
          </div>
        </div>
        <div className={styles.cardmini}>
          <h2>登録済みの体調</h2>
          <div className={styles.detail}>
          <ul>
              {allConditions.map((allCondition, index) => (
                <li key={index}>{allCondition.Name}</li>
              ))}
            </ul>
          </div>
          <button
            className={styles.ConditionButton}
            onClick={handleConditionClick}
          >
            体調の新規登録
          </button>
        </div>
      </div>
    </div>
  );
}
