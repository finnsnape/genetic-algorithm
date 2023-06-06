import { writable, type Writable } from 'svelte/store';
import type { GAStatus, CreateCategory } from './types';

export const gaStatus: Writable<GAStatus> = writable<GAStatus>({ fittestIndividual: undefined, generation: 0 });

export const createCategory: Writable<CreateCategory> = writable<CreateCategory>();