import { useState } from "react";
import type {  ReactNode } from "react";
import React from "react";

// T型: string型またはnumber型を表す共用型（Union Type）を定義
// type T = string | number;

// useFunc: カスタムフックの定義
// ジェネリック型パラメータTを指定し、extends unknownでJSXタグと区別
// [T, React.Dispatch<React.SetStateAction<T>>]の形で状態と更新関数を返す
export const useFunc = <T extends unknown>(initState: T): [T, React.Dispatch<React.SetStateAction<T>>] => {
    // useState: Reactのフック関数で状態管理を行う
    // TValue: 現在の状態値
    // setTvalue: 状態を更新するための関数
    const [TValue, setTvalue] = useState<T>(initState);
    
    // フックから状態と更新関数をタプルとして返す
    return [TValue, setTvalue];
};

// SubPageProps: SubPageコンポーネントのプロパティ型定義
type SubPageProps<T extends string | number> = {
    value: T;
};

// SubPage: 指定されたプロパティを表示するコンポーネント
// value: 表示する値
export function SubPage<T extends string | number>({ value }: SubPageProps<T>): ReactNode {
    // 受け取った値でカスタムフックを初期化
    const [TValue, _] = useFunc<T>(value);

    return (
        <>
            {value}
            <div>
                {TValue}
            </div>
        </>
    );
}