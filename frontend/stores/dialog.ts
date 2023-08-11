import { atom } from 'jotai'

export const commentDialogAtom = atom(false)
export const postDialogAtom = atom(false)
export const loginDialogAtom = atom(false)
export const collectionDialogAtom = atom<number | boolean>(false)
export const topicDialogAtom = atom(false)