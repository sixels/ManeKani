export interface CreateReviewRequest {
  /** The card ID */
  card: string;
  /** User errors */
  errors: { [key: string]: number };
  /** The review's study session type */
  session_type: "review" | "lesson";
}
