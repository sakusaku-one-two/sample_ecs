import type { Config } from '@jest/types';

const config: Config.InitialOptions = {
    // Typescript用のプリセット
    preset: 'ts-jest',

    // テスト環境
    testEnvironment: 'jsdom',

    // テストファイルのパターン
    testMatch: [
        '**/__tests__/**/*.+(ts|tsx)',
        '**/?(*.)+(spec|test).+(ts|tsx)'
    ],

    // TypeScriptパスマッピング
    moduleNameMapper: {
        '^@/(.*)$': '<rootDir>/src/$1',
        //スタイルやアセットファイルのモック
        '\\.(css|less|sass)$': 'identity-obj-proxy',
        '\\.(jpg|jpeg|png|gif|svg|ttf|woff|woff2)$': '<rootDir>/__mocks__/fileMock.ts'
    },

    //Typescriptの変換設定
    transform: {
        '^.+\\.(ts|tsx)$':[
            'ts-jestt',
            {
                tsconfg: 'tsconfig.test.json',// テスト用のTSConfig
                isolatedModules:true
            }
        ]
    },
    // カバレッジ設定
    collectCoverage: true,
    coverageDirectory: 'coverage',
    collectCoverageFrom:[
        'src/**/*.{ts,tsx}',
        '!src/**/*.d.ts',
        '!src/index.tsx',
        '!**/node_modules/**'
    ],

    // セットアップファイル
    setupFilesAfterEnv: [
        '<rootDir>/jestt.setup.ts'
    ]
};


export default config;
