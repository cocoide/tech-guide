import { PreviewDialogModel } from '@/types/dialog'
import { Article } from '@/types/model'
import { atom } from 'jotai'

export const postDialogAtom = atom(false)
export const loginDialogAtom = atom(false)
export const collectionDialogAtom = atom<number | boolean>(false)
export const topicDialogAtom = atom(false)
export const commentDialogAtom = atom<Article | boolean>(false)
export const previewDialogAtom = atom<PreviewDialogModel | boolean>(false)
