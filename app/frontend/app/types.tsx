type Book = {
  id: number;
  isbn: number;
  titleName: string;
  titleNameKana: string;
  publisherName: string;
  publisherNameKana: string;
  publishDate: string;
  price: number;
  status: 0 | 1 | 2 | 3;
  readingStartDate: string | null;
  readingEndDate: string | null;
};

type Author = {
  id: number;
  name: string;
  nameKana: string;
};

type Subject = {
  id: number;
  subjectName: string;
  subjectKana: string;
};

type BookInfo = {
  book: Book;
  authors: Author[];
  subjects: Subject[];
};
