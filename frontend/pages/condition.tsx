import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { format, isWithinInterval, parseISO, set } from "date-fns";
import styles from "./condition.module.scss";
import axios, { AxiosRequestConfig } from "axios";

// バックエンドからのレスポンス形式に合わせた型定義
type BackendCondition = {
  ID: number;
  Name: string;
  Type: string;
};

export default function ConditionRegistration() {
  const [cycleConditions, setCycleConditions] = useState<BackendCondition[]>(
    []
  );

  const [environmentConditions, setEnvironmentConditions] = useState<
    BackendCondition[]
  >([]);
  const [selectedCondition, setSelectedCondition] = useState<string | null>(
    null
  );
  const [errorMessage, setErrorMessage] = useState("");
  const [selectedConditionId, setSelectedConditionId] = useState<number | null>(
    null
  );
  const [startDate, setStartDate] = useState(format(new Date(), "yyyy-MM-dd"));
  const [duration, setDuration] = useState<number>(0);
  const [cycleLength, setCycleLength] = useState<number>(0);
  const [damage_point, setDamage_point] = useState<number>(0);
  const [region, setRegion] = useState<string>("Tokyo");

  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login");
    } else {
      fetchCycleConditions(token);
      fetchEnvironmentConditions(token);
    }
  }, [router]);
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_BASE_URL;
  const fetchCycleConditions = async (token: string) => {
    try {
      const response = await axios.get(
        `${backendUrl}/conditions/cycle`,
        {
          headers: {
            "x-token": token,
          },
        }
      );
      setCycleConditions(response.data);
    } catch (error) {
      console.error("Fetching cycle conditions failed", error);
      setErrorMessage("サイクル条件の取得に失敗しました。");
    }
  };

  const fetchEnvironmentConditions = async (token: string) => {
    try {
      const response = await axios.get(
        `${backendUrl}/conditions/environment`,
        {
          headers: {
            "x-token": token,
          },
        }
      );
      setEnvironmentConditions(response.data);
    } catch (error) {
      console.error("Fetching environment conditions failed", error);
      setErrorMessage("環境条件の取得に失敗しました。");
    }
  };

  const handleBackButtonClick = () => {
    router.push("/main");
  };

  const handleConditionSelection = (
    conditionId: number,
    conditionName: string
  ) => {
    setSelectedCondition(conditionName); // 条件名を選択された状態として設定
    setSelectedConditionId(conditionId); // 条件IDも選択された状態として設定
  };

  const handleCycleConditionSubmit = async () => {
    // ここでtokenを取得します
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("Token is not found");
      return;
    }

    try {
      const response = await axios.post(
        `${backendUrl}/users/me/condition/cycle`,
        {
          condition_id: selectedConditionId,
          start_date: startDate, // 開始日
          duration: duration, // 期間（日数）
          cycle_length: cycleLength, // 周期の長さ（日数）
          damage_point: damage_point, // ダメージポイント（初期値は0）
        },
        {
          headers: {
            "x-token": token,
          },
        }
      );
      setSelectedCondition(null);
      setSelectedConditionId(null);
      setStartDate(format(new Date(), "yyyy-MM-dd"));
      setDuration(0);
      setCycleLength(0);
      setDamage_point(0);

      // ここでレスポンスに基づいて適切な処理を行います
      console.log("Cycle Condition Created: ", response.data);
    } catch (error) {
      console.error("Error posting cycle condition", error);
    }
  };

  const handleEnvironmentConditionSubmit = async () => {
    // ここで環境条件の登録処理を実装します
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("Token is not found");
      return;
    }

    try {
      const response = await axios.post(
        `${backendUrl}/users/me/condition/environment`,
        {
          condition_id: selectedConditionId,
          start_date: startDate, // 開始日
          region: region, // 地域
          count: 0,
          damage_point: damage_point, // ダメージポイント（初期値は0）
        },
        {
          headers: {
            "x-token": token,
          },
        }
      );
      console.log({
        condition_id: selectedConditionId,
        start_date: startDate,
        region: region,
        count: 0,
        damage_point: damage_point,
      });

      setSelectedCondition(null);
      setSelectedConditionId(null);
      setStartDate(format(new Date(), "yyyy-MM-dd"));
      setRegion("Tokyo");
      setDamage_point(0);
    } catch (error) {
      console.error("Error posting enviroment condition", error);
    }
  };

  return (
    <div className={styles.conditionPage}>
      <div className={styles.conditionHeader}>
        <button
          className={styles.conditionButton}
          onClick={handleBackButtonClick}
        >
          mainに戻る
        </button>
      </div>
      <div className={styles.conditionContainer}>
        <div className={styles.card}>
          <div className={styles.cycleConditionButtons}>
            {cycleConditions.map((condition) => (
              <button
                key={condition.ID}
                onClick={() =>
                  handleConditionSelection(condition.ID, condition.Name)
                }
                className={
                  selectedConditionId === condition.ID
                    ? `${styles.conditionButton} ${styles.selectedConditionButton}` // 選択されたボタンのスタイルを適用
                    : styles.conditionButton
                }
              >
                {condition.Name}
              </button>
            ))}
          </div>
          <div className="cycleconditionfrom">
            <label className={styles.formLabel}>
              前回の開始日：
              <input
                type="date"
                value={startDate}
                onChange={(e) => setStartDate(e.target.value)}
              />
            </label>
            <label className={styles.formLabel}>
              期間：
              <input
                type="number"
                value={duration}
                onChange={(e) => setDuration(Number(e.target.value))}
              />
            </label>
            <label className={styles.formLabel}>
              平均サイクル：
              <input
                type="number"
                value={cycleLength}
                onChange={(e) => setCycleLength(Number(e.target.value))}
              />
            </label>
            <label className={styles.formLabel}>
              辛さ（数字で1〜100表してください）：
              <input
                type="number"
                min="1"
                max="100"
                value={damage_point}
                onChange={(e) => setDamage_point(Number(e.target.value))}
              />
            </label>
            <label className={styles.formLabel}>
              <button onClick={handleCycleConditionSubmit}>周期条件を登録</button>
            </label>
          </div>
        </div>

        <div className={styles.card}>
          <div className={styles.environmentConditionButtons}>
            {environmentConditions.map((condition) => (
              <button
                key={condition.ID}
                onClick={() =>
                  handleConditionSelection(condition.ID, condition.Name)
                }
                className={
                  selectedConditionId === condition.ID
                    ? `${styles.conditionButton} ${styles.selectedConditionButton}` // 選択されたボタンのスタイルを適用
                    : styles.conditionButton
                }
              >
                {condition.Name}
              </button>
            ))}

            <div className="enviromentconditionfrom">
              <label className={styles.formLabel}>
                居場所(ローマ字で、都道府県で入力)：
                <input
                  type="text"
                  value={region}
                  onChange={(e) => setRegion(e.target.value)}
                />
              </label>
              <label className={styles.formLabel}>
                辛さ（数字で1〜100表してください）：
                <input
                  type="number"
                  min="1"
                  max="100"
                  value={damage_point} // ここでは既に定義されている damage_point を使用
                  onChange={(e) => setDamage_point(Number(e.target.value))}
                />
              </label>
              <label className={styles.formLabel}>
                <button onClick={handleEnvironmentConditionSubmit}>
                  環境条件を登録
                </button>
              </label>
            </div>
          </div>
        </div>

        {errorMessage && <p className={styles.errorMessage}>{errorMessage}</p>}
      </div>
    </div>
  );
}
