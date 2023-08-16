infra/aws: AWS 관련 서비스 파일들을 추가하였습니다. (DynamoDB, S3, SES, Polly, FirebaseAuth 등)
infra/database: 데이터베이스 마이그레이션, 모델 정의 및 연결 설정 파일들을 추가하였습니다.
internal/app/handler & service: 각 거래소별 핸들러 및 서비스 로직 파일을 추가하였습니다. (binance를 예로 들었습니다.)
internal/domain: 도메인 모델을 정의한 파일들을 추가하였습니다.
internal/repository: 각 거래소별 데이터베이스 CRUD 작업을 처리할 파일을 추가하였습니다.
test: 유닛 테스트와 통합 테스트를 위한 디렉토리를 추가하였습니다.
api: OpenAPI 스펙을 관리할 yaml 파일을 추가하였습니다.
errors: 사용자 정의 에러 타입을 관리할 파일을 추가하였습니다.
.github/workflows: GitHub Actions을 위한 CI/CD 설정 파일을 추가하였습니다.


├─ cmd                             # 실행 가능한 애플리케이션의 진입점
│  ├─ asset-monitor                # 자산 모니터링 애플리케이션
│  │  └─ main.go                   # asset-monitor 애플리케이션의 시작점
│  └─ trade-log                    # 거래 기록 애플리케이션
│     └─ main.go                   # trade-log 애플리케이션의 시작점
├─ go.mod                          # 프로젝트의 의존성 모듈 관리
├─ go.sum                          # 의존성 모듈의 체크섬
├─ infra                           # 인프라 및 외부 서비스와의 통합 관련 코드
│  ├─ aws                          # AWS 서비스와 관련된 코드
│  │  ├─ config.go                 # AWS 설정 및 초기화
│  │  ├─ lambda.go                 # AWS Lambda 관련 함수
│  │  ├─ dynamodb.go               # DynamoDB 관련 코드
│  │  ├─ s3.go                     # S3 관련 코드
│  │  ├─ ses.go                    # SES 관련 코드
│  │  ├─ polly.go                  # Polly 관련 코드
│  │  └─ firebase_auth.go          # Firebase 인증 관련 코드
│  └─ database                     # 데이터베이스 관련 코드
│     ├─ migrations                # 데이터베이스 마이그레이션 스크립트
│     ├─ models.go                 # 데이터 모델 정의
│     └─ connection.go             # 데이터베이스 연결 및 설정
├─ internal                        # 애플리케이션의 핵심 비즈니스 로직
│  ├─ app                          # 애플리케이션 로직
│  │  ├─ handler                   # 요청 핸들러 (HTTP 요청 처리)
│  │  │  ├─ binance.go             # Binance 관련 핸들러
│  │  │  └─ ...                    # 다른 거래소 관련 핸들러
│  │  └─ service                   # 비즈니스 로직
│  │     ├─ binance.go             # Binance 관련 서비스 로직
│  │     └─ ...                    # 다른 거래소 관련 서비스 로직
│  ├─ domain                       # 도메인 모델 정의
│  │  ├─ exchange.go               # 거래소 도메인 모델
│  │  ├─ trade.go                  # 거래 도메인 모델
│  │  └─ balance.go                # 잔액 도메인 모델
│  └─ repository                   # 데이터베이스 CRUD 연산
│     ├─ binance.go                # Binance 관련 데이터베이스 연산
│     └─ ...                       # 다른 거래소 관련 데이터베이스 연산
├─ test                            # 테스트 코드
│  ├─ unit                         # 유닛 테스트 코드
│  └─ integration                  # 통합 테스트 코드
├─ api                             # API 문서 관련
│  └─ openapi_spec.yaml            # OpenAPI 스펙 문서
├─ errors                          # 에러 처리 관련 코드
│  └─ custom_errors.go             # 사용자 정의 에러 타입
├─ .github                         # GitHub 설정 및 워크플로우
│  └─ workflows                    # GitHub Actions 워크플로우
│     └─ ci_cd.yaml                # CI/CD 설정
└─ template.yml                    # AWS SAM 템플릿 (Lambda, API Gateway 설정 등)