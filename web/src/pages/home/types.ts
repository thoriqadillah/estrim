export interface UploadableFile {
  name: string;
  size: number;
  progress?: number;
}

export interface UploadedFile {
  id: string;
  user_id: string;
  job_id?: number;
  name: string;
  path: string;
  mime: string;
  type: string;
  size: number;
  is_compressed: boolean;
  created_at: Date;
}
